import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Optuznica } from 'src/app/model/sudstvo/optuznica';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class OptuznicaService {

  constructor(private client: HttpClient) { }

  getOptuznice() : Observable<Optuznica[]> {
    return this.client.get<Optuznica[]>(environment.apiUrl + "/tuzilastvo/optuznice")
  }

  // getOptuznica(id: string) : Observable<Optuznica> {
  //   return this.client.get<Optuznica>(environment.apiUrl + "/tuzilastvo/optuznice" + id);
  // }
}
