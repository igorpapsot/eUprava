import express from "express"
import { appController } from "./controllers/appController"

const app = express()

app.use(appController)

app.listen(8004, () => {
    console.log("MUP app started")
})