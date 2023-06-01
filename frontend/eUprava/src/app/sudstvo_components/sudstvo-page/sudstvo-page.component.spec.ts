import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SudstvoPageComponent } from './sudstvo-page.component';

describe('SudstvoPageComponent', () => {
  let component: SudstvoPageComponent;
  let fixture: ComponentFixture<SudstvoPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SudstvoPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SudstvoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
