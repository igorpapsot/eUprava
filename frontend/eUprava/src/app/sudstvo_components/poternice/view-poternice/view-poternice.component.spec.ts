import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPoterniceComponent } from './view-poternice.component';

describe('ViewPoterniceComponent', () => {
  let component: ViewPoterniceComponent;
  let fixture: ComponentFixture<ViewPoterniceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewPoterniceComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPoterniceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
