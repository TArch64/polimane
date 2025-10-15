import type { IHttpTransport } from './IHttpTransport';

const NETWORK_ERROR = new TypeError('Network request failed');
const TIMEOUT_ERROR = new TypeError('Network request timed out');
const ABORT_ERROR = new DOMException('Aborted', 'AbortError');

export class HttpLegacyTransport implements IHttpTransport {
  send(request: Request): Promise<Response> {
    return new Promise(async (resolve, reject) => {
      const xhr = new XMLHttpRequest();

      xhr.open(request.method, request.url);
      this.setRequestHeaders(xhr, request.headers);

      xhr.onload = () => resolve(this.buildResponse(xhr));
      xhr.onerror = () => reject(NETWORK_ERROR);
      xhr.ontimeout = () => reject(TIMEOUT_ERROR);
      xhr.onabort = () => reject(ABORT_ERROR);

      request.signal?.addEventListener('abort', () => {
        xhr.abort();
      });

      xhr.send(await this.getRequestBody(request));
    });
  }

  private setRequestHeaders(xhr: XMLHttpRequest, headers: Headers): void {
    for (const [key, value] of headers.entries()) {
      xhr.setRequestHeader(key, value);
    }
  }

  private buildResponse(xhr: XMLHttpRequest): Response {
    return new Response(xhr.responseText, {
      status: xhr.status,
      statusText: xhr.statusText,
      headers: this.buildResponseHeaders(xhr),
    });
  }

  private buildResponseHeaders(xhr: XMLHttpRequest): Headers {
    const entries = xhr.getAllResponseHeaders()
      .split('\r\n')
      .map((line) => {
        const [name, value] = line.split(': ');
        return [name, value] as [string, string];
      })
      .filter(([name, value]) => name && value);

    return new Headers(entries);
  }

  private async getRequestBody(request: Request): Promise<XMLHttpRequestBodyInit | null> {
    return request.body ? await request.clone().arrayBuffer() : null;
  }
}
