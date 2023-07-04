import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Sud } from 'src/app/model/sudstvo/sud';
import { SudEnum } from 'src/app/model/sudstvo/sudEnum';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class SudstvoService {

  constructor(private client: HttpClient) { }

  options() {
    return {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
  }

  postSud(sud : Sud) {
    console.log(sud)
    return this.client.post<unknown>(environment.apiUrl + "/sudstvo/sudovi", {
      sud : sud.sud,
      datum : sud.datum,
      mesto : sud.mesto,
    }, this.options())
  }

  getSudovi() : Observable<Sud[]> {
    return this.client.get<Sud[]>(environment.apiUrl + "/sudstvo/sudovi")
  }
  
  getSud(id: string) : Observable<Sud> {
    return this.client.get<Sud>(environment.apiUrl + "/sudstvo/sudovi" + id);
  }

}
