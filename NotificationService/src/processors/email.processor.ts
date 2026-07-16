import { Worker, Job } from "bullmq";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import { NotificationDto } from "../dto/notification.dto";
import { getRedisConnectionObject } from "../config/redis.config";


export const setupMailerWorker = new Worker<NotificationDto>(
  MAILER_QUEUE,
  async (job: Job) => {
    if (job.name != MAILER_PAYLOAD) {
      throw new Error("invalid job name");
    }
    const payload = job.data;
    console.log(`processing email for ${JSON.stringify(payload)}`);
  },
  {
    connection: getRedisConnectionObject(),
  },
);

setupMailerWorker.on("completed", (job) => {
  console.log(`job completed ${JSON.stringify(job)}`);
});

setupMailerWorker.on("failed", (job) => {
  console.log(`job failed ${JSON.stringify(job)}`);
});
