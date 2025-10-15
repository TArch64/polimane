import type { IHttpTransport } from './IHttpTransport';

export class HttpModernTransport implements IHttpTransport {
  send(request: Request): Promise<Response> {
    return fetch(request);
  }
}
