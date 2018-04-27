import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FlatBoardComponent } from './flat-board.component';

describe('FlatBoardComponent', () => {
  let component: FlatBoardComponent;
  let fixture: ComponentFixture<FlatBoardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FlatBoardComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FlatBoardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
