import { LoginType } from "../login-type";
import { RegisterType } from "../register-type";
import { ValidationError } from "./validation-error";
import { isBlank } from "@/core/utils/is-blank";

enum Dypesss {
  aaa = "aaa",
  bbb = "bbb",
}

export class HomePageValidation {
  //validationError = new ValidationError();
  error = new Map<string, string>();
  private blanksCounter = 0;
}
