const { v4: uuidv4 } = require("uuid");

export function generateIdempotencyKey(): string {
  return uuidv4();
}
