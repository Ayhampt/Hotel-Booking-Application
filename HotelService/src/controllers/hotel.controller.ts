import { NextFunction, Request, Response } from "express";
import { createHotelService, getAllHotelsService, getHotelByIdService, softDeleteHotelService, updateHotelService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";
export async function createHotelHandler(req:Request,res:Response,next:NextFunction) {

  const hotelResponse = await createHotelService(req.body)

  res.status(StatusCodes.CREATED).json({
    message:"hotel created successfully",
    data:hotelResponse,
    success:true

  })
}

export async function getHotelByIdHandler(req:Request,res:Response,next:NextFunction) {

  const hotelResponse = await getHotelByIdService(Number(req.params.id))

  res.status(StatusCodes.OK).json({
    message:"hotel found successfully",
    data:hotelResponse,
    success:true

  })
}

export async function getAllHotelsHandler(req:Request,res:Response,next:NextFunction) {
  const hotelResponse = await getAllHotelsService()

  res.status(StatusCodes.OK).json({
    message:"All hotels fetched successfully",
    data: hotelResponse,
    success:true
  })
}

export async function deleteHotelHandler(req:Request,res:Response,next:NextFunction) {
  const hotelResponse = await softDeleteHotelService(Number(req.params.id))
  res.status(StatusCodes.OK).json({
    message:"hotel deleted successfully",
    data:hotelResponse,
    success:true
  })
}

export async function updateHotelHandler(req:Request,res:Response,next:NextFunction) {
  const hotelResponse = await updateHotelService(Number(req.params.id),req.body)
  res.status(StatusCodes.OK).json({
    message:"hotel updated successfully",
    data:hotelResponse,
    success:true
  })
}