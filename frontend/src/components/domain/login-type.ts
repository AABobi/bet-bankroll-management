export type LoginType = {
  email: string;
  password: string;
};

function isLoginType(value: any): value is LoginType {
  return "email" in value && "password" in value;
}
