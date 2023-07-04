import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreatePoternicaComponent } from './create-poternica.component';

describe('CreatePoternicaComponent', () => {
  let component: CreatePoternicaComponent;
  let fixture: ComponentFixture<CreatePoternicaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreatePoternicaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreatePoternicaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
