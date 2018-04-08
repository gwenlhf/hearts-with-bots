import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { RulesService } from './rules.service';

import { AppComponent } from './app.component';
import { GameboardComponent } from './gameboard/gameboard.component';
import { GameCardComponent } from './game-card/game-card.component';
import { GameSidePipe } from './game-side.pipe';
import { UnicardPipe } from './unicard.pipe';


@NgModule({
  declarations: [
    AppComponent,
    GameboardComponent,
    GameCardComponent,
    GameSidePipe,
    UnicardPipe
  ],
  imports: [
    BrowserModule
  ],
  providers: [RulesService],
  bootstrap: [AppComponent]
})
export class AppModule { }
