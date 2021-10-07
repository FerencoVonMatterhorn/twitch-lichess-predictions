import {Component} from '@angular/core';
import {ApiService} from "./api.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'Twitch lichess predictions';

  constructor(private apiService: ApiService) {
    this.apiService.getLastGame().subscribe({
      next: data => {
        console.log(data)
      },
      error: error => {
        console.error('There was an error!', error);
      }
    })
  }
}
