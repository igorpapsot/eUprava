import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Optuznica } from 'src/app/model/tuzilastvo/optuznica';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class OptuzniceService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getOptuznice() : Observable<Optuznica[]> {
    return this.client.get<Optuznica[]>(environment.apiUrl + "/tuzilastvo/optuznice");
  }
}
