import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PoternicaComponent } from './poternica.component';

describe('PoternicaComponent', () => {
  let component: PoternicaComponent;
  let fixture: ComponentFixture<PoternicaComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PoternicaComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PoternicaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
