import { writable } from "svelte/store";

export type Account = {
  id: string;
  username: string;
  created_at: string
  updated_at: string;
}

export const account = writable<Account | null>(null);
