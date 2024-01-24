import { LoginType } from "./login-type";

export type RegisterType = {
  email: string;
  password: string;
  repeatedPassword: string;
};

function isRegisterType(value: any): value is RegisterType {
  return "username" in value && "password" in value;
}
