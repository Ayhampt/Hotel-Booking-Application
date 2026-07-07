/*
  Warnings:

  - You are about to drop the column `idemkey` on the `IdempotencyKey` table. All the data in the column will be lost.
  - A unique constraint covering the columns `[key]` on the table `IdempotencyKey` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `key` to the `IdempotencyKey` table without a default value. This is not possible if the table is not empty.

*/
-- DropIndex
DROP INDEX `IdempotencyKey_idemkey_key` ON `IdempotencyKey`;

-- AlterTable
ALTER TABLE `IdempotencyKey` DROP COLUMN `idemkey`,
    ADD COLUMN `key` VARCHAR(191) NOT NULL,
    MODIFY `id` INTEGER NOT NULL AUTO_INCREMENT;

-- CreateIndex
CREATE UNIQUE INDEX `IdempotencyKey_key_key` ON `IdempotencyKey`(`key`);
