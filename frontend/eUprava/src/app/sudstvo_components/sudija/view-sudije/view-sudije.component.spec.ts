import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewSudijeComponent } from './view-sudije.component';

describe('ViewSudijeComponent', () => {
  let component: ViewSudijeComponent;
  let fixture: ComponentFixture<ViewSudijeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewSudijeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewSudijeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
