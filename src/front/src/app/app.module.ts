import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { GameComponent } from './game/game.component';
import { GameCardComponent } from './game-card/game-card.component';
import { GamePlayerComponent } from './game-player/game-player.component';


@NgModule({
  declarations: [
    AppComponent,
    GameComponent,
    GameCardComponent,
    GamePlayerComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
