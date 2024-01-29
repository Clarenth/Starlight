// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';
import {context} from '../models';

export function CreateNewProfile(arg1:string,arg2:string):Promise<void>;

export function LoadAccountData(arg1:string):Promise<string>;

export function NewMemDB():Promise<{[key: string]: models.Account}>;

export function PrintDB():Promise<string>;

export function SetContext(arg1:context.Context):Promise<void>;
