import { z } from "zod";

export const roomGenerationRequestSchema = z.object({
  roomCategoryId: z.number().positive(),
  startDate: z.iso.datetime(),
  endDate: z.iso.datetime(),
  scheduleType: z.enum(["immediate", "scheduled"]).default("immediate"),
  scheduledAt: z.iso.datetime().optional(),
  priceOverride: z.number().positive(),
});

export const roomGenerationJobSchema = z.object({
  roomCategoryId: z.number().positive(),
  startDate: z.iso.datetime(),
  endDate: z.iso.datetime(),
  priceOverride: z.number().positive().optional(),
  batchSize: z.number().positive().default(100),
});

export type roomGenerationJob = z.infer<typeof roomGenerationJobSchema>;
export type roomGenerationRequest = z.infer<typeof roomGenerationRequestSchema>;


export interface RoomGenerationResponse {
  success: boolean;
  totalRoomsGenerated: number;
  totalDatesProcessed: number;
  errors: string[];
  jobId: string;
}
