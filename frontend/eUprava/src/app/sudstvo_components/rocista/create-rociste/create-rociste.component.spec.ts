import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRocisteComponent } from './create-rociste.component';

describe('CreateRocisteComponent', () => {
  let component: CreateRocisteComponent;
  let fixture: ComponentFixture<CreateRocisteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateRocisteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateRocisteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
