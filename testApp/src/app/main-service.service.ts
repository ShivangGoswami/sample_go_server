import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class MainServiceService {
  url = 'http://localhost:1234/endpoint'
  constructor(public http: HttpClient) { }

  getGreeting(){
    return this.http.get(this.url)
  }
}
