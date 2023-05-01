import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';
import { KrivicnaPrijava } from 'src/app/model/tuzilastvo/krivicnaPrijava';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class KrivicnaPrijavaServiceService {

  constructor(private client: HttpClient) { }

  options() {
    return  {
      headers: new HttpHeaders({
        'Content-Type':  'application/json',
        //'Authorization': `Bearer ${sessionStorage.getItem('token')}`,
      })
    };
  }

  getPrijave() : Observable<KrivicnaPrijava[]> {
    return this.client.get<KrivicnaPrijava[]>(environment.apiUrl + "/tuzilastvo/prijave");
  }

  declinePrijava(id: string) {
    return this.client.put<unknown>(environment.apiUrl + "/tuzilastvo/prijave/decline/" + id, null, this.options());
  }

  confirmPrijava(id: string) {
    return this.client.put<unknown>(environment.apiUrl + "/tuzilastvo/prijave/confirm/" + id, null, this.options());
  }
}