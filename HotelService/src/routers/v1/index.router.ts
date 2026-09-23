import express from "express";
import pingRouter from "./ping.router";
import hotelRouter from "./hotel.router";
import roomGenerationRouter from "./roomGeneration.router";

const v1Router = express();

v1Router.use("/ping", pingRouter);
v1Router.use("/hotels", hotelRouter);
v1Router.use("/room-generation",roomGenerationRouter) // Use require to import the router

export default v1Router;
