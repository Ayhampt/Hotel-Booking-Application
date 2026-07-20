import logger from "../config/logger";
import { serverConfig } from "../config";
import transporter from "../config/mailer.config";
import { InternalServerError } from "../utils/errors/app.error";

export const sendEmail = async (to: string, subject: string, body: string) => {
  try {
    await transporter.sendMail({
      from: serverConfig.MAILER_EMAIL,
      to,
      subject,
      html: body,
    });
    logger.info(`email send to ${to}`);
  } catch (error) {
    throw new InternalServerError("Cannot");
  }
};
