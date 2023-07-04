import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { KonacnaPresuda } from 'src/app/model/sudstvo/konacnaPresuda';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class KonacnaPresudaService {

  constructor(private client: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postKonacnaPresuda(konacnaPresuda : KonacnaPresuda) {
    console.log(konacnaPresuda)
    return this.client.post<unknown>(environment.apiUrl + "/sudstvo/konacnaPresuda",{
      aktivna : konacnaPresuda.aktivna,
      opis : konacnaPresuda.opis,
    }, this.options())
  }


  getKonacnaPresuda(id: string) : Observable<KonacnaPresuda> {
    return this.client.get<KonacnaPresuda>(environment.apiUrl + "/sudstvo/konacnePresuda" + id);
  }

}
