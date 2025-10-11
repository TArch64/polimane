export interface IHttpTransport {
  send(request: Request): Promise<Response>;
}
