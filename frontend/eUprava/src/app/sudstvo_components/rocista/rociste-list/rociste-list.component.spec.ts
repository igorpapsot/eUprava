import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RocisteListComponent } from './rociste-list.component';

describe('RocisteListComponent', () => {
  let component: RocisteListComponent;
  let fixture: ComponentFixture<RocisteListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RocisteListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RocisteListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
