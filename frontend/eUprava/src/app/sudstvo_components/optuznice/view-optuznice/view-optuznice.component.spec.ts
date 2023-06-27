import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewOptuzniceComponent } from './view-optuznice.component';

describe('ViewOptuzniceComponent', () => {
  let component: ViewOptuzniceComponent;
  let fixture: ComponentFixture<ViewOptuzniceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewOptuzniceComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewOptuzniceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
