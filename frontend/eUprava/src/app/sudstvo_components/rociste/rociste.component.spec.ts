import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RocisteComponent } from './rociste.component';

describe('RocisteComponent', () => {
  let component: RocisteComponent;
  let fixture: ComponentFixture<RocisteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RocisteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RocisteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
