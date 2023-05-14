import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Sud } from 'src/app/model/sudstvo/sud';
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

  getSudovi() : Observable<Sud[]> {
    return this.client.get<Sud[]>(environment.apiUrl + "sudstvo/sudovi")
  }

}
