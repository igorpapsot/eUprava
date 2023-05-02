import express from "express"
import { UserController } from "./controllers/user.controller"
import mongoose from "mongoose"
import bodyParser from "body-parser"

async function main() {

    const app = express()
    console.log("Connecting to db...")
    const db = await mongoose.connect("mongodb://localhost:27018/mup")
    console.log("Connection established.")

    app.use(bodyParser.urlencoded({extended: true}))
    app.use(bodyParser.json())
    app.use(UserController)

    app.listen(8004, () => {
        console.log("MUP app started")
    })
}
main();