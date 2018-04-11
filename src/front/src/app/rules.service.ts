import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';

import { Game } from './game';

import { SAMPLE_TEXT, SAMPLE_OUT } from './sample';

@Injectable()
export class RulesService {
	getNewGame() : Observable<Game> {
		return of(SAMPLE_TEXT);
	}
	getNewGameFlat() : Observable<number[]> {
		return of(SAMPLE_OUT);
	}
  constructor() { }
}
