import { type IDBPDatabase, openDB } from 'idb';

export type IdbStorageKey<V extends object> = string & { __type: V };

interface IdbStorageConfig {
  name: string;
  version: number;
  store: string;
}

export class IdbStorage {
  private static _instance: IdbStorage;

  static get instance(): IdbStorage {
    this._instance ??= new IdbStorage({
      name: 'polimane',
      version: 1,
      store: 'keyval',
    });

    return this._instance;
  }

  static key<V extends object>(key: string): IdbStorageKey<V> {
    return key as IdbStorageKey<V>;
  }

  private constructor(private readonly config: IdbStorageConfig) {}

  private database?: Promise<IDBPDatabase>;

  private accessDatabase(): Promise<IDBPDatabase> {
    this.database ??= openDB(this.config.name, this.config.version, {
      upgrade: (db: IDBPDatabase) => {
        db.createObjectStore(this.config.store);
      },
    });

    return this.database;
  }

  async getItem<V extends object>(key: IdbStorageKey<V>): Promise<V | null> {
    const db = await this.accessDatabase();
    return (await db.get(this.config.store, key)) ?? null;
  }

  async setItem<V extends object>(key: IdbStorageKey<V>, value: V): Promise<void> {
    const db = await this.accessDatabase();
    await db.put(this.config.store, value, key);
  }
}
