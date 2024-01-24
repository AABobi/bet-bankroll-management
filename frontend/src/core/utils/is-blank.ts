export function isBlank(value: string | Date | number | undefined): boolean {
  if (value === undefined) {
    return true;
  }

  if (typeof value === "string") {
    return value.trim() === "";
  }
  if (typeof value === "number") {
    return isNaN(value);
  }
  if (value instanceof Date) {
    return isNaN(value.getTime());
  }

  return false;
}
