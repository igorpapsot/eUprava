import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { Tuzilastvo } from 'src/app/model/tuzilastvo/tuzilastvo';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TuzilastvoService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getTuzilastva() : Observable<Tuzilastvo[]> {
    return this.client.get<Tuzilastvo[]>(environment.apiUrl + "/tuzilastvo/tuzilastva");
  }
}
