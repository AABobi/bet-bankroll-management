import { LoginType } from "../login-type";
import { RegisterType } from "../register-type";
import { isBlank } from "@/core/utils/is-blank";

export class ValidationError {
  error = new Map<string, string>();

  addError(field: string, error: string) {
    this.error.set(field, error);
  }

  removeError(field: string) {
    this.error.delete(field);
  }

  getError(field: string): string | undefined {
    return this.error.get(field);
  }

  reportBlank(value: LoginType | RegisterType): boolean {
    let wasBlank = 0;
    for (const key in value) {
      if (isBlank((value as any)[key])) {
        this.addError(key, "Cannot be blank");
        wasBlank++;
      }
    }
    if (wasBlank !== 0) {
      return true;
    }
    return false;
  }
}
