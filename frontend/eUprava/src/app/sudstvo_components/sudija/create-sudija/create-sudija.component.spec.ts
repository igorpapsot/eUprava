import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateSudijaComponent } from './create-sudija.component';

describe('CreateSudijaComponent', () => {
  let component: CreateSudijaComponent;
  let fixture: ComponentFixture<CreateSudijaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateSudijaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateSudijaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
