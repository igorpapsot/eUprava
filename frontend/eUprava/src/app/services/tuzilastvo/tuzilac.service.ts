import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { Tuzilac } from 'src/app/model/tuzilastvo/tuzilac';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TuzilacService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getTuzilac(jmbg: string) : Observable<Tuzilac> {
    return this.client.get<Tuzilac>(environment.apiUrl + "/tuzilastvo/tuzioci/" + jmbg);
  }
}
