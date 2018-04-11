import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';

import { Game } from './game';

import { SAMPLE_TEXT } from './sample';

@Injectable()
export class RulesService {
	getNewGame() : Observable<Game> {
		return of(SAMPLE_TEXT);
	}
  constructor() { }
}
