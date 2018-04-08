import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RulesService } from './rules.service';

import { AppComponent } from './app.component';
import { GameboardComponent } from './gameboard/gameboard.component';
import { GameCardComponent } from './game-card/game-card.component';


@NgModule({
  declarations: [
    AppComponent,
    GameboardComponent,
    GameCardComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [RulesService],
  bootstrap: [AppComponent]
})
export class AppModule { }
