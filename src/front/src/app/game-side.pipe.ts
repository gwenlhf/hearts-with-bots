import { Pipe, PipeTransform } from '@angular/core';

import { Side } from './game.constants';

@Pipe({
  name: 'gameSide'
})
export class GameSidePipe implements PipeTransform {

  transform(value: Side): string {
    return Side[value];
  }

}
