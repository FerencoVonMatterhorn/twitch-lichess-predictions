import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { LichessComponent } from './lichess/lichess.component';
import { TwitchComponent } from './twitch/twitch.component'
import { RouterModule, Routes} from "@angular/router";
import { HomeComponent } from './home/home.component';
import { DrowdownDirective } from "./shared/drowdown.directive";

const appRoutes: Routes =[
  { path: 'twitch', component: TwitchComponent },
  { path: 'home', component: HomeComponent },
  { path: 'lichess', component: LichessComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LichessComponent,
    TwitchComponent,
    HomeComponent,
    DrowdownDirective
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
