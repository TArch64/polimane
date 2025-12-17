export const authChannel = new BroadcastChannel('auth-channel');

export enum AuthChannelEvent {
  COMPLETE = 'auth-complete',
  DELETED_USER = 'auth-deleted-user',
}
