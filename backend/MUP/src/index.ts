import express from "express"
import { UserController } from "./controllers/user.controller"
import mongoose from "mongoose"
import bodyParser from "body-parser"
import { ReqController } from "./controllers/req.controller"
import cors from "cors"

async function main() {

    const app = express()
    console.log("Connecting to db...")
    const db = await mongoose.connect("mongodb://mup_db:27017/mup")
    console.log("Connection established.")

    app.use(cors())
    app.use(bodyParser.urlencoded({extended: true}))
    app.use(bodyParser.json())
    
    app.use(UserController)
    app.use(ReqController)

    app.listen(8004, () => {
        console.log("MUP app started")
    })
}
main();