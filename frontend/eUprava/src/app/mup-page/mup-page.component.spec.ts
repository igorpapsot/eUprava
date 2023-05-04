import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MupPageComponent } from './mup-page.component';

describe('MupPageComponent', () => {
  let component: MupPageComponent;
  let fixture: ComponentFixture<MupPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MupPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MupPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
