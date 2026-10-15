public class ShoeManufacture {

    public static Shoe produceShoe(String shoeType)
    {
        ShoeFactory shoeFactory = null;
        if ("Formal".equals(shoeType))
        {
            shoeFactory = new FormalShoeFactory();
        }
        else if ("Sports".equals(shoeType))
        {
            shoeFactory = new SportsShoeFactory();
        }
        else if ("Casual".equals(shoeType))
        {
            shoeFactory = new CasualShoeFactory();
        }
        return new Shoe(shoeFactory.createShoeSole(), shoeFactory.createShoeLace());
    }
    public static void main(String[] args) {
        Shoe formalShoe = produceShoe("Formal");
        Shoe sportaShoe = produceShoe("Sports");
        Shoe casualShoe = produceShoe("Casual");

        formalShoe.displayBuildShoe();
        sportaShoe.displayBuildShoe();
        casualShoe.displayBuildShoe();



    }
    
}
