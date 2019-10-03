
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';


@Injectable({
  providedIn: 'root'
})

export class MovieService {
  private baseUrl: string = environment.api;

  constructor(private http: HttpClient){
  }

  public getMovies(id?: number){
    const header = new HttpHeaders({'Content-Type': 'application/json'});
    return this.http.post<any>(this.baseUrl, {"ID": id}, {headers: header})
  }
}
