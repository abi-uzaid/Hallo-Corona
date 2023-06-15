import { Container, Col, Row, Card } from "react-bootstrap";
import { useQuery } from "react-query";
import { Link } from "react-router-dom";
import { API } from "../../config/api";

function Article() {
  let { data: articles } = useQuery("articlessCache", async () => {
    const response = await API.get("/articles");
    // console.log(response);
    return response.data.data;
  });

  return (
    <Container>
      {articles?.length !== 0 ? (
        <Row>
          {articles?.map((item, index) => (
            <Col md={3} sm={6} xs={12} className="mb-5" key={index}>
              <Card className="h-100 shadow-lg">
                <div>
                  <Card.Img
                    variant="top"
                    style={{ width: "260px", height: "215px" }}
                    src={item.image}
                  />
                </div>
                <Link
                  to={`/article/${item.id}`}
                  className="text-decoration-none mb-2"
                  style={{ color: "black", fontSize:"15px" }}
                >
                  <Card.Body>
                    <Card.Title className="text-truncate" >{item.title}</Card.Title>
                    <Card.Text
                      className="text-truncate"
                      style={{ color: "#6C6C6C" }}
                    >
                      {item.desc}
                    </Card.Text>
                  </Card.Body>
                </Link>
                <p className="category px-3 py-1 rounded-pill ms-3 mb-3">
                  {item.category}
                </p>
              </Card>
            </Col>
          ))}
        </Row>
      ) : (
        ""
      )}
    </Container>
  );
}
export default Article;
