import type { HttpErrorReason } from './HttpErrorReason';

export interface IHttpCommonErrorResponse {
  data: { reason: HttpErrorReason };
}

export class HttpError extends Error {
  static async fromResponse(response: Response): Promise<HttpError> {
    const headers = response.headers ?? new Headers();
    const isJson = headers.get('content-type')?.includes('application/json');
    const body = isJson ? await response.json() : await response.text();
    return new HttpError(body);
  }

  constructor(
    readonly response: unknown,
  ) {
    super('HTTP error');
  }

  get reason(): HttpErrorReason | null {
    return (this.response as IHttpCommonErrorResponse)?.data?.reason ?? null;
  }
}
