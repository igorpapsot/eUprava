import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewRocisteComponent } from './view-rociste.component';

describe('ViewRocisteComponent', () => {
  let component: ViewRocisteComponent;
  let fixture: ComponentFixture<ViewRocisteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewRocisteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewRocisteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
