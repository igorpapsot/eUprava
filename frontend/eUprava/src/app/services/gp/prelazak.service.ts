import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {environment} from "src/environments/environment";
import {Observable} from "rxjs";
import {ProveraGradjanina} from "src/app/model/GP/proveraGradjanina";
import {PrelazakRequest} from "src/app/model/GP/prelazakRequest";
import {PrelazakGranice} from "../../model/GP/prelazakGranice";

@Injectable({
  providedIn: 'root'
})
export class PrelazakService {

  constructor(private httpClient: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }


  postPrelazak(prelazak : PrelazakRequest) {
    return  this.httpClient.post<unknown>(environment.apiUrl + "/gp/prelazak", {
      policajacId : prelazak.policajacId,
      proveraId: prelazak.proveraId
    }, this.options())
  }

  getPrelaske() : Observable<PrelazakGranice[]> {
    return this.httpClient.get<PrelazakGranice[]>(environment.apiUrl + "/gp/prelasci")
  }

}

