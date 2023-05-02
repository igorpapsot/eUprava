import { Router } from "express";

export const appController = Router();

appController.get("/app", async (req, res) => {
    res.json({
        name: "test"
    })
    return
})