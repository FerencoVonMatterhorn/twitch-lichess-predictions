import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private httpClient: HttpClient) {
  }

  public getLastGame() {
    const headers = {'Content-Type': 'application/json', 'Accept': 'application/json'}
    return this.httpClient.get("https://lichess.org/api/user/ferenco/current-game", {headers});
  }
}
