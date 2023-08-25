import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {ProveraRequest} from "src/app/model/GP/proveraRequest";
import {environment} from "src/environments/environment";
import {Observable} from "rxjs";
import {ProveraGradjanina} from "../../model/GP/proveraGradjanina";

@Injectable({
  providedIn: 'root'
})
export class ProveraService {

  constructor(private httpClient: HttpClient) {}

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postProvera(provera : ProveraRequest) {
    console.log(provera)
    return this.httpClient.post<unknown>(environment.apiUrl + "/gp/provera", {
      policajacId : provera.PolicajacId,
      gradjanin : provera.Gradjanin
    }, this.options())
  }

  getProvere() : Observable<ProveraGradjanina[]> {
    return this.httpClient.get<ProveraGradjanina[]>(environment.apiUrl + "/gp/provere")
  }

  getProvereNaCekanju() : Observable<ProveraGradjanina[]> {
    return this.httpClient.get<ProveraGradjanina[]>(environment.apiUrl + "/gp/provere/cekaju")
  }

  postPusti(provera : string) {
    console.log(provera)
    return this.httpClient.post<unknown>(environment.apiUrl + "/gp/provera/accept/" + provera, {
    }, this.options())
  }

  postZabrani(provera : string) {
    console.log(provera)
    return this.httpClient.post<unknown>(environment.apiUrl + "/gp/provera/accept/" + provera, {
    }, this.options())
  }

}
