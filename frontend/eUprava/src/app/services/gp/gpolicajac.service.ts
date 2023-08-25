import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {environment} from "src/environments/environment";
import {Jwt} from "src/app/model/GP/jwt";

@Injectable({
  providedIn: 'root'
})
export class GpolicajacService {

  constructor(private httpClient: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postLoginPolicajca(jmbg : string, password : string) {
    console.log(jmbg + password)
    return this.httpClient.post<Jwt>(environment.apiUrl + "/gp/provera", {
      jmbg : jmbg,
      password : password
    }, this.options())
  }
}
