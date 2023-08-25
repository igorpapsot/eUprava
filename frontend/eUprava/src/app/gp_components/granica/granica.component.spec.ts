import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GranicaComponent } from './granica.component';

describe('GranicaComponent', () => {
  let component: GranicaComponent;
  let fixture: ComponentFixture<GranicaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GranicaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GranicaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
